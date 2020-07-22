import request from '@/utils/request'

// 查询在线用户列表
export function list(query) {
  return request({
    url: '/monitor/online',
    method: 'get',
    params: query
  })
}

// 强退用户
export function forceLogout(token) {
  return request({
    url: '/monitor/online',
    method: 'delete',
    params: {
      tokenStr: token
    }
  })
}
