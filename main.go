package main

import (
    "os"

    "git.solver4all.com/Peter/kp"
    "git.solver4all.com/Peter/s4a"
    "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()
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
	    Action: babSolver,
        },
	{
	    Name: "hs",
	    Usage: "Solve knapsack problem by branch and bound algorithm of Horowitz and Sahni",
	    Action: hsSolver,
        },
	{
	    Name: "dp",
	    Usage: "Solve knapsack problem by dynamic programming",
	    Action: dpSolver,
	},
	{
	    Name: "greedy",
	    Usage: "Solve knapsack problem by greedy heuristic",
	    Action: greedySolver,
	},
	{
	    Name: "ub",
	    Usage: "Compute an upper bound for the objective function value of a knapsack problem",
	    Action: upperBound,
	},
	{
	    Name: "gen",
	    Usage: "Generate a knapsack problem instance",
	    Action: generate,
	},
    }

    app.Run(os.Args)
}

func babSolver(c *cli.Context) error {
    var (
        kpp kp.KnapsackProblemT
	err error
    )

    // read
    err = readData(&kpp, c)
    if err != nil {
        return err
    }

    //check
    err = kp.CheckEqualLength(&kpp)
    if err != nil {
        return err
    }
    err = kp.CheckSortedItems(&kpp)
    if err != nil {
        return err
    }

    // solve
    kp.BranchAndBound(&kpp)

    // write
    return writeKnapsackProblem(&kpp, c)
}

func hsSolver(c *cli.Context) error {
    var (
        kpp kp.KnapsackProblemT
	err error
    )

    // read
    err = readData(&kpp, c)
    if err != nil {
        return err
    }

    //check
    err = kp.CheckEqualLength(&kpp)
    if err != nil {
        return err
    }
    err = kp.CheckSortedItems(&kpp)
    if err != nil {
        return err
    }

    // solve
    kp.BranchAndBound(&kpp)

    // write
    return writeKnapsackProblem(&kpp, c)
}

func dpSolver(c * cli.Context) error {
    var (
        kpp kp.KnapsackProblemT
	err error
    )

    // read
    err = readData(&kpp, c)
    if err != nil {
        return err
    }

    // check
    err = kp.CheckEqualLength(&kpp)
    if err != nil {
        return err
    }

    // solve
    kp.DynProg(&kpp)

    // write
    return writeKnapsackProblem(&kpp, c)
}

func greedySolver(c *cli.Context) error {
    var (
        kpp kp.KnapsackProblemT
	err error
    )

    // read
    err = readData(&kpp, c)
    if err != nil {
        return err
    }

    //check
    err = kp.CheckEqualLength(&kpp)
    if err != nil {
        return err
    }
    err = kp.CheckSortedItems(&kpp)
    if err != nil {
        return err
    }

    // solve
    kp.Greedy(&kpp)

    // write
    return writeKnapsackProblem(&kpp, c)
}

func upperBound(c *cli.Context) error {
    var (
        kpp kp.KnapsackProblemT
	err error
    )

    // read
    err = readData(&kpp, c)
    if err != nil {
        return err
    }

    //check
    err = kp.CheckEqualLength(&kpp)
    if err != nil {
        return err
    }
    err = kp.CheckSortedItems(&kpp)
    if err != nil {
        return err
    }

    // solve
    kp.UpperBound(&kpp)

    // write
    return writeKnapsackProblem(&kpp, c)
}

func generate(c *cli.Context) error {
    var (
        kpgen kp.KnapsackGeneratorT
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
    return writeKnapsackProblem(kpp, c)
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

func writeKnapsackProblem(kpp *kp.KnapsackProblemT, c *cli.Context) error {
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
