/*
Copyright 2017 wechat-go Authors. All Rights Reserved.
MIT License

Copyright (c) 2017
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
SOFTWARE.
*/

package wxweb

import (
	"time"
	"github.com/skip2/go-qrcode"
)

const ()

var ()

// CreateSession: create wechat bot session
// if common is nil, session will be created with default config
// if handlerRegister is nil,  session will create a new HandlerRegister
func T_CreateSession() (*Session, []byte, error) {
	common := DefaultCommon
	
	wxWebXcg := &XmlConfig{}
	
	// get qrcode
	uuid, err := JsLogin(common)
	if err != nil {
		return nil, nil, err
	}
	
	session := &Session{
		WxWebCommon: common,
		WxWebXcg:    wxWebXcg,
		QrcodeUUID:  uuid,
		CreateTime:  time.Now().Unix(),
	}
	
	session.HandlerRegister = CreateHandlerRegister()
	
	png, err := qrcode.Encode("https://login.weixin.qq.com/l/"+uuid, qrcode.Medium, 256)
	if err != nil {
		return nil, nil, err
	}
	return session, png, nil
}
