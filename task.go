package workflow

// Task task is runnable, i parameter for custom type conversions
type Task interface {
	Run(ctx context.Context, i interface{}) error
}
