/*
 * File Created: Tuesday, 20th October 2020 2:30:51 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */
import request from '@/helpers/request';

export function generateAnalytics() {
    return request({
        url: `/analytics-job/generate-now`,
        method: "POST",
        withCredentials: true
    })
}