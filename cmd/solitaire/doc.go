/*
 *  Copyright 2025 Markus Mahlberg <138420+mwmahlberg@users.noreply.github.com>
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

/*
Package main provides a command line utility for the Solitaire encryption algorithm.
It provides a simple way to encrypt and decrypt files using the Solitaire algorithm.
The tool is only intended for small messages and should not be used beyond a small amount of data,
a couple of thousane characters at most.
The algorithm is not suitable for large files or long messages.
It uses a playing card deck to generate a keystream, which is then used to encrypt and decrypt the data.

	$ go install github.com/mwmahlberg/solitaire/cmd/solitaire
	$ solitaire -h
*/
package main
