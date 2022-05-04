import { ElNotification } from 'element-plus';


/**
 * 成功弹框
 */
export const $notificatioSuccess = (message = '成功') => {
  return ElNotification({
    title: '成功',
    message,
    type: 'success',
  })
}

/**
 * 异常弹框
 */
export const $notificatioError = (message = '失败') => {
  return ElNotification({
    title: '失败',
    message,
    type: 'error',
  })
}

/**
 * 警告弹框
 */
 export const $notificatioWarn = (message = '警告') => {
  return ElNotification({
    title: '警告',
    message,
    type: 'warning',
  })
}

/**
 * 解析json
 * @param data 
 * @returns 
 */
export const $parse = data => {
  let res = {}
  try {
    res = JSON.parse(data)
  } catch (error) {
  }
  return res
}

export const $toJson = data => {
  let res = ''
  try {
    res = JSON.stringify(data)
  } catch (error) {
  }
  return res
}