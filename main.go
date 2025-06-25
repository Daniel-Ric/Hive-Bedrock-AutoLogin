package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/auth"
)

func main() {
	region := flag.String("region", "", "Region to connect: na, eu, asia")
	duration := flag.Duration("duration", 2*time.Minute, "Connection stay time")
	seed := flag.Int64("seed", time.Now().UnixNano(), "Random seed for delays")
	flag.Parse()

	if *region == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Select region:")
		fmt.Println("1) North America (na)")
		fmt.Println("2) Europe        (eu)")
		fmt.Println("3) Asia          (asia)")
		fmt.Print("Choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		switch input {
		case "1", "na":
			*region = "na"
		case "2", "eu":
			*region = "eu"
		case "3", "asia":
			*region = "asia"
		default:
			color.Yellow("Unknown choice, defaulting to Europe (eu)")
			*region = "eu"
		}
	}

	var server string
	switch *region {
	case "na":
		server = "ca.hivebedrock.network:19132"
	case "eu":
		server = "fr.hivebedrock.network:19132"
	case "asia":
		server = "sg.hivebedrock.network:19132"
	default:
		server = "fr.hivebedrock.network:19132"
	}

	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgHiBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Printf("%s Server: %s (%s)\n", blue("[INFO]"), server, *region)
	fmt.Println(blue("[AUTH] Starting Device Flow authentication..."))
	tokenSrc := auth.WriterTokenSource(os.Stdout)
	fmt.Println(blue("[AUTH] Please complete login in browser to proceed."))
	if _, err := tokenSrc.Token(); err != nil {
		fmt.Printf("%s %v\n", red("[ERROR] Authentication failed:"), err)
		os.Exit(1)
	}
	fmt.Println(green("[SUCCESS] Authenticated successfully."))

	rand.Seed(*seed)
	fmt.Println(blue("[INFO] Entering main loop."))

	dialer := minecraft.Dialer{TokenSource: tokenSrc}

	for {
		var conn *minecraft.Conn
		for {
			fmt.Printf("%s Attempting to connect to %s...\n", blue("[CONNECT]"), server)
			c, err := dialer.DialContext(context.Background(), "raknet", server)
			if err != nil {
				fmt.Printf("%s Connection failed: %v\n", red("[ERROR]"), err)
				fmt.Println(blue("[RETRY] Retrying in 2m..."))
				time.Sleep(2 * time.Minute)
				continue
			}
			conn = c
			fmt.Println(green("[SUCCESS] Connected successfully."))
			break
		}

		fmt.Printf("%s Staying connected for %s...\n", green("[INFO]"), duration)
		time.Sleep(*duration)

		fmt.Println(yellow("[DISCONNECT] Disconnecting..."))
		if err := conn.Close(); err != nil {
			fmt.Printf("%s Disconnect error: %v\n", red("[ERROR]"), err)
		} else {
			fmt.Println(green("[SUCCESS] Disconnected cleanly."))
		}

		minDelay := 23 * time.Hour
		extraMin := rand.Intn(61)
		wait := minDelay + time.Duration(extraMin)*time.Minute
		next := time.Now().Add(wait)

		fmt.Printf("%s Next connection scheduled at %s\n", blue("[NEXT]"), next.Format("2006-01-02 15:04:05"))

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = fmt.Sprintf(" Waiting %s until next attempt", wait)
		s.Color("cyan")
		s.Start()
		time.Sleep(wait)
		s.Stop()
		fmt.Println()
	}
}
