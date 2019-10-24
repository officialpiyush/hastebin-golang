/**
 * Copyright (C) 2019 Piyush Bhangale
 *
 * This file is part of hastebin-golang.
 *
 * hastebin-golang is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * hastebin-golang is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with hastebin-golang.  If not, see <http://www.gnu.org/licenses/>.
 */

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var config map[string]string

func Load() bool {
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
		return false
	}

	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func Get(c string) string {
	return config[c]
}
