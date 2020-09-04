import request from '@/helpers/request';

export function isExistsCanvasConfig() {
    return request({
        url: `/settings/canvas`,
        method: "GET",
        withCredentials: true
    })
}

export function storeOrUpdateCanvasConfig(formData){
    return request({
        url: `/settings/canvas`,
        method: "POST",
        data: formData,
        withCredentials: true
    })
}

export function settings(){
    return request({
        url: `/settings`,
        method: "GET",
        withCredentials: true
    })
}