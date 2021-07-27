import axios from 'axios'
axios.defaults.withCredentials = true
axios.defaults.xsrfHeaderName = 'X-CSRFToken'
axios.defaults.xsrfCookieName = 'csrftoken'

// ====================================================自己的django ==============================
const instanceFoo = axios.create({
    baseURL: '/foo',
    timeout: 8000
})

/*
 *封装get方法
 *@param{String} url [请求地址]
 *@param{Object} params 请求参数
 */
export function Get2 (url, params) {
    return new Promise((resolve, reject) => {
        instanceFoo.get(url, {
            params: params
        }).then((res) => {
            resolve(res.data);
        }).catch((error) => {
            reject(error.data);
        })
    })
}

/**
 *封装post方法
 *@param{String} url 请求地址
 *@param{Object} params 请求参数
 */
export function Post2 (url, params) {
    return new Promise((resolve, reject) => {
        instanceFoo.post(url, params).then((res) => {
            resolve(res.data);
        }).catch((error) => {
            reject(error.data);
        })
    })
}

// ===================================== oj-backend ===================================
const instance = axios.create({
    baseURL: '/api',
    timeout: 10000,
})

// 请求拦截
instance.interceptors.request.use(config => {
    return config;
}, error => {
    return Promise.reject(error);
})

// 响应拦截
instance.interceptors.response.use(response => {

    const resCode = response.status;
    if (resCode === 200) {
        if (response.data.code === 401) {
            window.location.href = '/';
        }
        return Promise.resolve(response);
    } else {
        return Promise.resolve(response);
    }
}, error => {
    // const resCode = error.response.status;
    // switch (resCode) {
    //   case 401:
    //     // vm.$Message.error(error.response.data.message);
    //     // store.commit('logout', this);
    //     // store.commit('clearOpenedSubmenu');
    //     // this.selfLog('token-0', store.state.app.token);
    //     window.location.href = '/'
    //     break;
    //   case 404:
    //     vm.$Message.error('网络请求不存在');
    //     break;
    //   case 500:
    //     vm.$Message.error('服务器连接错误');
    //     break;
    //   // 其他状态码错误提示
    //   default:
    //     vm.$Message.error(error.response.data.message);
    // }
    // return Promise.reject(error);
})

/*
 *封装get方法
 *@param{String} url [请求地址]
 *@param{Object} params 请求参数
 */
export function Get (url, params) {
    return new Promise((resolve, reject) => {
        instance.get(url, {
            params: params
        }).then((res) => {
            resolve(res.data);
        }).catch((error) => {
            reject(error.data);
        })
    })
}

/**
 *封装post方法
 *@param{String} url 请求地址
 *@param{Object} params 请求参数
 */
export function Post (url, params) {
    return new Promise((resolve, reject) => {
        instance.post(url, params).then((res) => {
            resolve(res.data);
        }).catch((error) => {
            reject(error.data);
        })
    })
}

/**
 *封装put方法
 *@param{String} url 请求地址
 *@param{Object} params 请求参数
 */
export function Put (url, params) {
    return new Promise((resolve, reject) => {
        instance.put(url, params).then((res) => {
            resolve(res.data);
        }).catch((error) => {
            reject(error.data);
        })
    })
}

/**
 *封装patch方法
 *@param{String} url 请求地址
 *@param{Object} params 请求参数
 */
export function Patch (url, params) {
    return new Promise((resolve, reject) => {
        instance.put(url, params).then((res) => {
            resolve(res.data);
        }).catch((error) => {
            reject(error.data);
        })
    })
}

/**
 *封装delete方法
 *@param{String} url [请求地址]
 *@param{Object} params [请求参数]
 */
export function Delete (url, params) {
    return new Promise((resolve, reject) => {
        instance.delete(url, {
            params: params
        }).then((res) => {
            resolve(res.data);
        }).catch((error) => {
            reject(error.data);
        })
    })
}
// ====================================================飞书的API================================
const instanceFS = axios.create({
    baseURL: '/open-apis',
    timeout: 8000
})
/*
 *封装get方法
 *@param{String} url [请求地址]
 *@param{Object} params 请求参数
 */
export function Get4 (url, params, headers) {
    return new Promise((resolve, reject) => {
        instanceFS.get(url, {
            params: params,
            headers: headers
        }).then((res) => {
            resolve(res.data);
        }).catch((error) => {
            reject(error.data);
        })
    })
}

/**
 *封装post方法
 *@param{String} url 请求地址
 *@param{Object} params 请求参数
 */
export function Post4 (url, params, headers) {
    return new Promise((resolve, reject) => {
        if (headers !== "") {
            instanceFS.defaults.headers = headers
        }
        instanceFS.post(url, params).then((res) => {
            resolve(res.data);
        }).catch((error) => {
            reject(error.data);
        })
    })
}
// ====================================================成电飞书API=============================
const instanceUestc_api = axios.create({
    baseURL: '/uestcapi',
    timeout: 8000
})
/*
 *封装get方法
 *@param{String} url [请求地址]
 *@param{Object} params 请求参数
 */
export function Get5 (url, params) {
    return new Promise((resolve, reject) => {
        instanceUestc_api.get(url, {
            params: params
        }).then((res) => {
            resolve(res.data);
        }).catch((error) => {
            reject(error.data);
        })
    })
}

/**
 *封装post方法
 *@param{String} url 请求地址
 *@param{Object} params 请求参数
 */
export function Post5 (url, params) {
    return new Promise((resolve, reject) => {
        instanceUestc_api.post(url, params).then((res) => {
            resolve(res.data);
        }).catch((error) => {
            reject(error.data);
        })
    })
}
// ====================================================学校的API================================
const instanceCkd = axios.create({
    baseURL: '/ckd',
    timeout: 4000
})

/*
 *封装get方法
 *@param{String} url [请求地址]
 *@param{Object} params 请求参数
 */
export function Get3 (url, params) {
    return new Promise((resolve, reject) => {
        instanceCkd.get(url, {
            params: params
        }).then((res) => {
            resolve(res.data);
        }).catch((error) => {
            reject(error.data);
        })
    })
}

/**
 *封装post方法
 *@param{String} url 请求地址
 *@param{Object} params 请求参数
 */
export function Post3 (url, params) {
    return new Promise((resolve, reject) => {
        instanceCkd.post(url, params).then((res) => {
            resolve(res.data);
        }).catch((error) => {
            reject(error.data);
        })
    })
}


// =========================================== cookie ====================================
//获取cookie、
export function getCookie (name) {
    var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
    if (arr = document.cookie.match(reg)) {
        return (arr[2]);
    }
    else
        return null;
}

//设置cookie,增加到vue实例方便全局调用
export function setCookie (c_name, value, expiredays) {
    var exdate = new Date();
    exdate.setDate(exdate.getDate() + expiredays);
    document.cookie = c_name + "=" + escape(value) + ((expiredays == null) ? "" : ";expires=" + exdate.toGMTString());
};

//删除cookie
export function delCookie (name) {
    var exp = new Date();
    exp.setTime(exp.getTime() - 1);
    var cval = getCookie(name);
    if (cval != null)
        document.cookie = name + "=" + cval + ";expires=" + exp.toGMTString();
};

