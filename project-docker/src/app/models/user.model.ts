export interface User {
    userInfos : {
        crossfitlovID: number;
        firstname:string;
        age: number;
        sexe: ["Boy","Girl"];
        email:string;
        boxCity:string;
        biography:string;
        job:string;
        gender:string;
        createAt:string;
    };
    tokenInfos : {
        name: string;
        value: string;
        expiresAt : string;
    };
}