import { User } from "./user.model";

export interface SigninData {
    userData : User,
    tokenData : {
        name: string
        value: string
        expiresAt : string
    } 
}
