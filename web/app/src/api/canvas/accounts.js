import request from '@/helpers/request';

export function listAccount(accountID) {
    return request({
        url: `/canvas/accounts/${accountID}/sub_accounts`,
        method: "GET",
        withCredentials: true
    })
}