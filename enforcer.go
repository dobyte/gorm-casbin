package casbin

import (
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm/logger"
	"time"
)

type Enforcer = casbin.Enforcer
type Logger = logger.Interface

type Options struct {
	Model    string        `json:"model"`    // model config file path
	Debug    bool          `json:"debug"`    // debug mode
	Enable   bool          `json:"enable"`   // enable permission
	Autoload bool          `json:"autoload"` // auto load policy
	Duration time.Duration `json:"duration"` // auto load duration
	Database interface{}   `json:"database"` // database instance, Choose between DB and Link parameters. If DB exists, use DB first.
	Table    string        `json:"table"`    // database policy table name
	Logger   Logger        `json:"logger"`   // database logger interface
}

func NewEnforcer(opts *Options) (*Enforcer, error) {
	adp, err := newAdapter(opts.Database, opts.Table, opts.Logger)
	if err != nil {
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer(opts.Model, adp)
	if err != nil {
		return nil, err
	}

	enforcer.EnableLog(opts.Debug)
	enforcer.EnableEnforce(opts.Enable)
	enforcer.EnableAutoNotifyWatcher(opts.Autoload)
	enforcer.EnableAutoSave(true)

	return enforcer, nil
}
