import { HeaderCreatePost } from './headerCreatePost'
import { Stories } from './stories'
import { ShowPosts } from './showPosts'
import { useEffect, useState } from 'react'
import type { UserModel } from '@/models/UserModel'
import { getUserData } from '@/services/api/user/getUserData'

export const Main = () => {

  const [userData, setUserData] = useState<UserModel>()
  
    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const data = await getUserData("1")
                setUserData(data ?? undefined)
            } catch (error) {
                console.error('Error fetching user data:', error)
            }
        }
        fetchUserData()
    }, [])

  return (
    <main className="flex flex-col gap-6 p-4 bg-gray-100">

      <HeaderCreatePost userData={userData} />

      <Stories userData={userData} />

      <ShowPosts />
      
    </main>
  )
}