package compressor

import "context"

func Start(w *Worker, ctx context.Context, runID string) { w.Ramp(ctx, runID, 12) }
