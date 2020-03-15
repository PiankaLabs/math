/*
 *   Copyright 2020 Pianka Labs LLC (https://www.piankalabs.com)
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package main

import (
	// System
	"fmt"
	"log"
	"os"
	"strings"

	// Third Party
	"github.com/urfave/cli/v2"

	// Internal
	eval "piankalabs.com/math/pkg/evaluator"
)

func main() {
	app := &cli.App{
		Name:  "math",
		Usage: "Evaluate a mathematical expression",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				expression := strings.Join(c.Args().Slice(), " ")
				evaluator := new(eval.Evaluator)
				tree, err := evaluator.Evaluate(expression)

				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(tree.Calculate())
			} else {
				err := cli.ShowAppHelp(c)

				if err != nil {
					log.Fatal(err)
				}
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
