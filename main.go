/*
Copyright © 2020 Alexander Kiryukhin <a.kiryukhin@mail.ru>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"context"
	"log"

	"github.com/neonxp/rutina/v2"

	"github.com/neonxp/chatcloud/pkg/config"
	"github.com/neonxp/chatcloud/pkg/db"
	"github.com/neonxp/chatcloud/pkg/server"
)

func main() {
	r := rutina.New(rutina.Opt.SetListenOsSignals(true))
	cfg, err := config.New()
	if err != nil {
		log.Println(err)
		return
	}
	database, err := db.New(cfg.MongoConnection, cfg.MongoName)
	if err != nil {
		log.Println(err)
		return
	}
	api := server.NewServer(database, cfg)
	r.Go(api.Run, rutina.RunOpt.SetOnDone(rutina.Shutdown))
	r.Go(func(ctx context.Context) error {
		<-ctx.Done()
		if err := api.Close(); err != nil {
			log.Println(err)
		}
		return nil
	}, nil)
	if err := r.Wait(); err != nil {
		log.Println(err)
	}
}
