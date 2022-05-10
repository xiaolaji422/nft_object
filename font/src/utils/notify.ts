/**
 * 统一处理通知和弹窗
 */
import { ElNotification } from 'element-plus';
import { ElMessageBox } from 'element-plus';
import { nextTick } from 'vue';
const notifyMessage = async (info, title, type) => {
    await nextTick();
    ElNotification({
        title: title,
        message: info,
        type: type,
        duration: 1000,
        position: 'top-right',
        offset: 75
    });
};

export const notify = {
    success: (info, title = '成功') => notifyMessage(info, title, 'success'),
    error: (info, title = '失败') => notifyMessage(info, title, 'error'),
    warning: (info, title = '温馨提示') => notifyMessage(info, title, 'warning'),
    info: (info, title = '提示') => notifyMessage(info, title, 'info')
};

const messageConfirm = (text, title, type, options: any = { confirmButtonText: '确定', cancelButtonText: '取消' }) => {
    return new Promise((resolve) => {
        ElMessageBox.confirm(text, title, {
            confirmButtonText: options.confirmButtonText,
            cancelButtonText: options.cancelButtonText,
            type: type
        })
            .then(() => {
                resolve(true);
            })
            .catch(() => {
                resolve(false);
            });
    });
};

export const confirm = {
    success: (info, title = '成功') => messageConfirm(info, title, 'success'),
    error: (info, title = '错误') => messageConfirm(info, title, 'error'),
    warning: (info, title = '温馨提示') => messageConfirm(info, title, 'warning'),
    info: (info, title = '提示') => messageConfirm(info, title, 'info')
};
