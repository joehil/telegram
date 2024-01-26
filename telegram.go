/*MIT License
Copyright (c) 2021 Joerg Hillebrand
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE. */

package main

import (
        "fmt"
        "os"
        "strings"
        tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
        _ "github.com/lib/pq"
        "github.com/spf13/viper"
)

var tbtoken string
var chatid  int64

func main() {
        //      viper.SetConfigName("yourCryptoBot") // name of config file (without extension)
        viper.AddConfigPath("/etc/") // path to look for the config file in

        viper.SetConfigName("telegram") // name of config file (name of exchange)

        // Read config
        read_config()

        sendTelegram(strings.Join(os.Args[1:], " "))
}

func sendTelegram(message string) {

        bot, err := tgbotapi.NewBotAPI(tbtoken)
        if err != nil {
                fmt.Printf("Telegram error: %v\n", err)
                return
        }

//      fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

        msg := tgbotapi.NewMessage(chatid, message)
        bot.Send(msg)
}

func read_config() {
        err := viper.ReadInConfig() // Find and read the config file
        if err != nil {             // Handle errors reading the config file
                fmt.Printf("Config file not found: %v\n", err)
        }

        tbtoken = viper.GetString("tbtoken")
        chatid  = int64(viper.GetInt("chatid"))
//      fmt.Printf("%d\n", chatid)
}
