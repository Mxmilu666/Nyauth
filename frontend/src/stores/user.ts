import { reactive } from 'vue'
import { defineStore } from 'pinia'

export interface UserInfo {
  is_banned: boolean
  register_at: string
  role: string
  user_avatar: string
  user_email: string
  user_id: string
  user_uuid: string
  user_name: string
}

export const useUserStore = defineStore('user', () => {
  const userInfo = reactive<UserInfo>({
    is_banned: false,
    register_at: '',
    role: '',
    user_avatar: '',
    user_email: '',
    user_id: '',
    user_uuid: '',
    user_name: ''
  })
  
  const updateUserInfo = (info: Partial<UserInfo>) => {
    Object.assign(userInfo, info)
  }

  return {
    userInfo,
    updateUserInfo
  }
})
