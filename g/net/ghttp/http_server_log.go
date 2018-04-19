// Copyright 2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.
// 默认错误日志封装.

package ghttp

import (
    "fmt"
    "strings"
    "gitee.com/johng/gf/g/util/gconv"
)

// 处理服务错误信息，主要是panic，http请求的status由access log进行管理
func (s *Server) handleAccessLog(r *Request) {
    if !s.IsAccessLogEnabled() {
        return
    }
    // 自定义错误处理
    if v := s.GetLogHandler(); v != nil {
        v(r)
        return
    }
    status := gconv.String(r.Response.status)
    if v := r.Response.Header().Get("Status Code"); v != "" {
        status = v
    }
    content := fmt.Sprintf(`%s %s %s`,     r.Method, r.URL.String(), r.Proto)
    content += fmt.Sprintf(", host: %s",   r.Host)
    content += fmt.Sprintf(", from: %s",   strings.Split(r.RemoteAddr, ":")[0])
    content += fmt.Sprintf(", refer: %s",  r.Referer())
    content += fmt.Sprintf(", status: %v", status)
    s.logger.Info(content)
}

// 处理服务错误信息，主要是panic，http请求的status由access log进行管理
func (s *Server) handleErrorLog(error interface{}, r *Request) {
    if !s.IsErrorLogEnabled() {
        return
    }
    // 自定义错误处理
    if v := s.GetLogHandler(); v != nil {
        v(r, error)
        return
    }
    content := fmt.Sprintf(`%s %s %s`,    r.Method, r.URL.String(), r.Proto)
    content += fmt.Sprintf(", host: %s",  r.Host)
    content += fmt.Sprintf(", from: %s",  strings.Split(r.RemoteAddr, ":")[0])
    content += fmt.Sprintf(", refer: %s", r.Referer())
    content += fmt.Sprintf(", error: %v", error)
    s.logger.Error(content)
}