import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, Router } from '@angular/router';
import { Observable } from 'rxjs';
import { AuthService } from '../services/auth.services';
import { User } from '../models/user.model';

@Injectable({
    providedIn: 'root'
})


export class AuthGuard implements CanActivate {

    private user: User;

    constructor(private auth: AuthService, private router: Router) {
        this.auth.user.subscribe(data => this.user = data)
    }
    canActivate(
        route: ActivatedRouteSnapshot,
        state: RouterStateSnapshot
    ): Observable<boolean> | Promise<boolean> | boolean {
        // if(this.user == null) {
        //     this.router.navigate(['/login']);
        //     return false;
        // }
        
        return true;
        
    }
}
