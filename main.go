package main

import (
    "os"
    "time"

    "git.solver4all.com/Peter/kp/kp"
    "git.solver4all.com/Peter/s4a"
    "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()
    app.Version = "0.8.15"
    app.Compiled = time.Now()
    app.Authors = []cli.Author{
        cli.Author{
	    Name: "Peter Becker",
	    Email: "peter@solver4all.com",
	},
    }
    app.Copyright = "(c) 2017 Peter Becker"
    app.Usage = "Knapsack problem solver and generator."
    app.Flags = []cli.Flag{
        cli.StringFlag{
	    Name: "input,i",
	    Value: "",
	    Usage: "read solver input data from file instead of standard input",
	},
	cli.StringFlag{
	    Name: "output,o",
	    Value: "",
	    Usage: "write result to file instead of standard output",
	},
    }
    app.Commands = []cli.Command{
	{
	    Name: "bab",
	    Usage: "Solve knapsack problem by branch and bound (A*)",
	    Action: func(c *cli.Context) error {
	        return solve(c, func(p kp.KnapsackProblem) ([]int,int) { return kp.BranchAndBound(p) })
	    },
        },
	{
	    Name: "hs",
	    Usage: "Solve knapsack problem by branch and bound algorithm of Horowitz and Sahni",
	    Action: func(c *cli.Context) error {
	        return solve(c, func(p kp.KnapsackProblem) ([]int,int) { return kp.BranchAndBoundHS(p) })
	    },
        },
	{
	    Name: "dp",
	    Usage: "Solve knapsack problem by dynamic programming",
	    Action: func(c *cli.Context) error {
	        return solve(c, func(p kp.KnapsackProblem) ([]int,int) { return kp.DynProg(p) })
	    },
	},
	{
	    Name: "greedy",
	    Usage: "Solve knapsack problem by greedy heuristic",
	    Action: func(c *cli.Context) error {
	        return solve(c, func(p kp.KnapsackProblem) ([]int,int) { return kp.Greedy(p) })
	    },
	},
	{
	    Name: "ub",
	    Usage: "Compute an upper bound for the objective function value of a knapsack problem",
	    Action: func(c *cli.Context) error {
	        return bound(c, func(p kp.KnapsackProblem) ([]float64,int) { return kp.UpperBound(p) })
	    },
	},
	{
	    Name: "gen",
	    Usage: "Generate a knapsack problem instance",
	    Action: generate,
	},
    }

    app.Run(os.Args)
}

func solve(c *cli.Context, solvfunc func(p kp.KnapsackProblem) ([]int,int)) error {
    var (
        kpp kp.KnapsackData
	err error
    )

    err = readData(&kpp, c)		// read
    if err != nil {
        return err
    }

    /*
    err = kp.CheckEqualLength(&kpp)	//check
    if err != nil {
        return err
    }
    */
    err = kp.CheckSortedItems(&kpp)
    if err != nil {
        return err
    }

    x,z := solvfunc(kpp)		// solve
    kpp.X = x
    kpp.Z = z

    return writeKnapsackProblem(&kpp, c)	// write

}

func bound(c *cli.Context, ubfunc func(p kp.KnapsackProblem) ([]float64,int)) error {
    var (
        kpp kp.KnapsackData
	err error
    )

    err = readData(&kpp, c)		// read
    if err != nil {
        return err
    }

    /*
    err = kp.CheckEqualLength(&kpp)	//check
    if err != nil {
        return err
    }
    */
    err = kp.CheckSortedItems(&kpp)
    if err != nil {
        return err
    }

    x,z := ubfunc(kpp)			// solve
    kpp.Xf = x
    kpp.Z = z

    return writeKnapsackProblem(&kpp, c)	// write
}

func generate(c *cli.Context) error {
    var (
        kpgen kp.KnapsackGenData
	err   error
    )

    // read
    err = readData(&kpgen, c)
    if err != nil {
        return err
    }

    // generate
    kpp, err := kp.Generate(kpgen)
    if err != nil {
        return err
    }

    // write
    return writeKnapsackProblem(&kpp, c)
}

func readData(object interface{}, c *cli.Context) error {
    var (
	r   *os.File
	err error
    )

    if ifn:=c.GlobalString("input"); ifn != "" {
        r, err = os.Open(ifn)
	if err != nil {
	    return err
	}
	defer r.Close()
    } else {
        r = os.Stdin
    }

    err = s4a.ReadFJsonInput(r, object)
    if err != nil {
        return err
    }

    return nil
}

func writeKnapsackProblem(kpp *kp.KnapsackData, c *cli.Context) error {
    var (
	w   *os.File
        err error
    )

    if ofn:=c.GlobalString("output"); ofn != "" {
        w, err = os.Create(ofn)
	if err != nil {
	    return err
	}
	defer w.Close()
    } else {
        w = os.Stdout
    }

    err = s4a.WriteFJsonOutput(w, &kpp)
    if err != nil {
        return  err
    }

    return nil
}
