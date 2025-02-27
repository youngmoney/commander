package main

import ()

func Match(name string, path string, commands *[]Command) *Command {
	for _, c := range *commands {
		if c.Name != name {
			continue
		}
		if c.MatchPathRegex != nil && !c.MatchPathRegex.MatchString(path) {
			continue
		}
		if c.MatchCommand != "" {
			if err := ExecuteCommandQuietly(c.MatchCommand, []string{}); err != nil {
				continue
			}
		}

		return &c
	}
	return nil
}
