package fireworksai

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/instill-ai/component/base"
)

func TestComponent_Execute(t *testing.T) {
	c := qt.New(t)

	bc := base.Component{}
	cmp := Init(bc)

	c.Run("ok - supported chat task", func(c *qt.C) {
		_, err := cmp.CreateExecution(base.ComponentExecution{
			Component: cmp,
			Task:      TaskTextGenerationChat,
		})
		c.Assert(err, qt.IsNil)

	})
	c.Run("ok - supported embedding task", func(c *qt.C) {
		_, err := cmp.CreateExecution(base.ComponentExecution{
			Component: cmp,
			Task:      TaskTextEmbeddings,
		})
		c.Assert(err, qt.IsNil)
	})

	c.Run("nok - unsupported task", func(c *qt.C) {
		_, err := cmp.CreateExecution(base.ComponentExecution{
			Component: cmp,
			Task:      "FOO",
		})
		c.Check(err, qt.ErrorMatches, "unsupported task")
	})
}