export interface UserInfos {
    crossfitlovID: number;
    firstname:string;
    age: number;
    email:string;
    boxCity:string;
    biography:string;
    job:string;
    gender:string;
    createdAt:string;
}


export interface User {
    userInfos : UserInfos;
    tokenInfos : {
        name: string;
        value: string;
        expiresAt : string;
    };
}