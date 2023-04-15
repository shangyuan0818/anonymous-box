package repo

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/database/dal"
	"github.com/star-horizon/anonymous-box-saas/database/model"
	"github.com/star-horizon/anonymous-box-saas/pkg/cache"
)

type SettingRepo interface {
	GetByName(ctx context.Context, name string) (string, error)
	GetIntByName(ctx context.Context, name string) (int, error)
	GetBoolByName(ctx context.Context, name string) (bool, error)

	ListByNames(ctx context.Context, names []string) (map[string]string, error)

	SetByName(ctx context.Context, name, value string) error
}

type settingRepo struct {
	fx.In
	Query *dal.Query
	Cache cache.Driver
}

func NewSettingRepo(repo settingRepo) SettingRepo {
	return &repo
}

// GetByName implements SettingRepo.GetByName.
func (r *settingRepo) GetByName(ctx context.Context, name string) (string, error) {
	ctx, span := tracer.Start(ctx, "get-setting-by-name", trace.WithAttributes(
		attribute.String("name", name),
	))
	defer span.End()

	if v, exist := r.Cache.Get(ctx, fmt.Sprint("database:setting:", name)); exist {
		if setting, ok := v.(string); ok {
			span.AddEvent("get-from-cache", trace.WithAttributes(
				attribute.String("status", "hit"),
			))

			return setting, nil
		}
	} else {
		span.AddEvent("get-from-cache", trace.WithAttributes(
			attribute.String("status", "miss"),
		))
	}

	setting, err := r.Query.Setting.WithContext(ctx).Where(r.Query.Setting.Name.Eq(name)).First()
	if err != nil {
		span.RecordError(err)
		return "", err
	}

	if err := r.Cache.Set(ctx, fmt.Sprint("database:setting:", name), setting.Value, 0); err != nil {
		span.RecordError(err)
	}

	return setting.Value, nil
}

// GetIntByName implements SettingRepo.GetIntByName.
func (r *settingRepo) GetIntByName(ctx context.Context, name string) (int, error) {
	ctx, span := tracer.Start(ctx, "get-setting-int-by-name")
	defer span.End()

	setting, err := r.GetByName(ctx, name)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(setting)
}

// GetBoolByName implements SettingRepo.GetBoolByName.
func (r *settingRepo) GetBoolByName(ctx context.Context, name string) (bool, error) {
	ctx, span := tracer.Start(ctx, "get-setting-bool-by-name")
	defer span.End()

	setting, err := r.GetByName(ctx, name)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(setting)
}

// ListByNames implements SettingRepo.ListByNames.
func (r *settingRepo) ListByNames(ctx context.Context, names []string) (map[string]string, error) {
	ctx, span := tracer.Start(ctx, "list-settings-by-names", trace.WithAttributes(
		attribute.StringSlice("names", names),
	))
	defer span.End()

	data, missed, err := r.Cache.GetMulti(ctx, lo.Map(names, func(name string, _ int) string {
		return fmt.Sprint("database:setting:", name)
	}))
	if err != nil {
		return nil, err
	}

	var settings = make(map[string]string, len(names))
	for key := range data {
		if setting, ok := data[key].(string); ok {
			settings[strings.TrimPrefix(key, "database:setting:")] = setting
		} else {
			missed = append(missed, key)
		}
	}

	span.AddEvent("get-multi-from-cache", trace.WithAttributes(
		attribute.StringSlice("missed", missed),
		attribute.StringSlice("hit", lo.Keys(settings)),
	))

	if len(missed) > 0 {
		slice, err := r.Query.Setting.WithContext(ctx).Where(r.Query.Setting.Name.In(names...)).Find()
		if err != nil {
			span.RecordError(err)
			return nil, err
		}

		settings = lo.Assign(settings, lo.SliceToMap(slice, func(setting *model.Setting) (string, string) {
			return setting.Name, setting.Value
		}))

		// cache settings
		if err := r.Cache.SetMulti(ctx, lo.MapEntries(settings, func(key string, value string) (string, any) {
			return key, value
		}), "database:setting:"); err != nil {
			span.RecordError(err)
		}

		span.AddEvent("set-multi-to-cache", trace.WithAttributes(
			attribute.StringSlice("keys", lo.Keys(settings)),
		))
	}

	return settings, nil
}

// SetByName implements SettingRepo.SetByName.
func (r *settingRepo) SetByName(ctx context.Context, name, value string) error {
	ctx, span := tracer.Start(ctx, "set-setting-by-name", trace.WithAttributes(
		attribute.String("name", name),
		attribute.String("value", value),
	))
	defer span.End()

	if err := r.Cache.Delete(ctx, fmt.Sprint("database:setting:", name)); err != nil {
		span.RecordError(err)
		return err
	}

	if _, err := r.Query.Setting.WithContext(ctx).
		Where(r.Query.Setting.Name.Eq(name)).
		Update(r.Query.Setting.Name, value); err != nil {
		span.RecordError(err)
		return err
	}

	if err := r.Cache.Set(ctx, fmt.Sprint("database:setting:", name), value, 0); err != nil {
		span.RecordError(err)
	}

	return nil
}
