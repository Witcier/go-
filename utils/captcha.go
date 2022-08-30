package utils

import (
	"context"
	"time"
	"witcier/go-api/global"

	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type CaptchaRedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func NewCaptchaRedisStore() *CaptchaRedisStore {
	return &CaptchaRedisStore{
		Expiration: time.Second * global.Config.Captcha.Expires,
		PreKey:     global.Config.Captcha.PrefixKey,
	}
}

func (crs *CaptchaRedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	crs.Context = ctx
	return crs
}

func (crs *CaptchaRedisStore) Set(id, value string) error {
	err := global.Redis.Set(crs.Context, crs.PreKey+id, value, crs.Expiration).Err()
	if err != nil {
		global.Log.Error("CaptchaRedisStore Set Error", zap.Error(err))
		return err
	}

	return nil
}

func (crs *CaptchaRedisStore) Get(key string, clear bool) string {
	val, err := global.Redis.Get(crs.Context, key).Result()
	if err != nil {
		global.Log.Error("CaptchaRedisStore Get Error", zap.Error(err))
		return ""
	}

	if clear {
		err := global.Redis.Del(crs.Context, key).Err()
		if err != nil {
			global.Log.Error("CaptchaRedisStore Del Error", zap.Error(err))
			return ""
		}
	}

	return val
}

func (crs *CaptchaRedisStore) Verify(id, answer string, clear bool) bool {
	key := crs.PreKey + id
	v := crs.Get(key, clear)

	return v == answer
}
