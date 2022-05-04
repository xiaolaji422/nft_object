/*
 * @Description  : 页面缓存处理
 */
export const session = {
  set(key, value) {
    if (typeof value === 'undefined') return false

    // obj为undefined或null或空字符串不能存储，布尔值可存储，但取时为字符串
    if (!value && (typeof value === 'undefined' || typeof value === 'object' || typeof value === 'string')) return false

    let saveStr = '';
    if (typeof value === 'object') {
      saveStr = JSON.stringify(value);
    } else {
      saveStr = value;
    }
    sessionStorage.setItem(key, saveStr);
    return true;
  },
  get(key) {
    if (typeof key !== 'string') {
      return;
    }
    let savedStr = sessionStorage.getItem(key);
    // 非法值返回，包括undefined、null、空字符串
    if (!savedStr) return undefined;

    let result;
    if ((!savedStr.includes('"') && !savedStr.includes('\\')) || !savedStr.includes(':')) {
      return savedStr;
    }
    try {
      result = JSON.parse(savedStr);
      return result;
    } catch {
      return savedStr;
    }
  },
  remove(key) {
    if (typeof key !== 'string') return;
    sessionStorage.removeItem(key);
  },
  clear(key) {
    if (key) return;
    sessionStorage.clear();
  }
};

export const local = {
  set(key, value) {
    if (typeof value === 'undefined') return false;

    // obj为undefined或null或空字符串不能存储，布尔值可存储，但取时为字符串
    if (!value && (typeof value === 'undefined' || typeof value === 'object' || typeof value === 'string')) return false;

    let saveStr = '';
    if (typeof value === 'object') {
      saveStr = JSON.stringify(value);
    } else {
      saveStr = value;
    }
    localStorage.setItem(key, saveStr);
    return true;
  },
  get(key) {
    if (typeof key !== 'string') {
      return;
    }
    let savedStr = localStorage.getItem(key);
    // 非法值返回，包括undefined、null、空字符串
    if (!savedStr) return;

    let result;
    if ((!savedStr.includes('"') && !savedStr.includes('\\')) || !savedStr.includes(':')) {
      return savedStr;
    }
    try {
      result = JSON.parse(savedStr);
      return result;
    } catch {
      return savedStr;
    }
  },
  remove(key) {
    if (typeof key !== 'string') return;
    //
    localStorage.removeItem(key);
  },
  clear(key) {
    if (key) return;
    localStorage.clear();
  }
};
