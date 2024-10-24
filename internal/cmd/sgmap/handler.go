// Code generated by kubectl-plugin-builder.

/* MIT License
 *
 * Copyright (c) 2024 naka-gawa
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package sgmap

import (
	"fmt"

	"github.com/naka-gawa/kubectl-sgmap/internal/cmd"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// this assignment ensures
// options struct must implement CLINodeOptions interface.
var _ cmd.CLINodeOptions = &options{}

type options struct {
	cmd        *cobra.Command
	args       []string
	streams    *genericclioptions.IOStreams
	outputMode cmd.OutputMode
}

// Complete implements CLINodeOptions interface.
func (o *options) Complete(cmd *cobra.Command, args []string) error {
	o.cmd = cmd
	o.args = args
	o.outputMode = sgmapOutputModeFlag
	return nil
}

// Validate implements CLINodeOptions interface.
func (o *options) Validate() error {
	return nil
}

// Run implements CLINodeOptions interface.
func (o *options) Run() error {
	switch o.outputMode {
	// case cmd.OutputModeJSON:
	// case cmd.OutputModeYAML:
	case cmd.OutputModeNormal:
		_, err := fmt.Fprintf(o.streams.Out, "%s\n", o.cmd.Use)
		return err
	}

	return fmt.Errorf("unsupported output format '%s' found", o.outputMode)
}
