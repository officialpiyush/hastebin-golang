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

package main

import (
	"hastebin-golang/config"
	"hastebin-golang/database"
	"hastebin-golang/router"
)

func main() {
	config.Load()
	database.InitiateClient()
	router.SetupRouter()
}
