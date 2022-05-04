const dataMap = {
    // 基本状态
    status: [
        { text: "启用", value: 1 },
        { text: "禁用", value: 0 },
    ],
    sex: [
        { text: "男", value: 1 },
        { text: "女", value: 0 },
    ],

    // faq 相关
    show_os: [
        { text: "Android小程序", value: 24 },
        { text: "Android", value: 14 },
        { text: "Win", value: 48 },
        { text: "IOS小程序", value: 25 },
        { text: "IOS", value: 15 },
    ],
    //  faq 段落类型
    faq_text_type: [
        { value: 1, text: "标题&内容" },
        { value: 2, text: "QA问答" },
        { value: 3, text: "图文组合" },
        { value: 4, text: "温馨提示" },
    ],
    product_msdk: [
        { value: 3, text: "V3版本" },
        { value: 5, text: "V5版本" },
    ],
    //  游戏发行
    product_issuer: [
        { value: 1, text: "腾讯" },
        { value: 2, text: "景秀" },
    ],



};

export function getMapData(properties) {
    if (dataMap && dataMap[properties]) return dataMap[properties];
    return null;
}