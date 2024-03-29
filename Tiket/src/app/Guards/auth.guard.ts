import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, UrlTree, Router } from '@angular/router';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {

  constructor(private router: Router){}

  canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
      if(this.loggedIn())return true;

    this.router.navigate(['/'])
  }
  
  loggedIn():boolean{
    return(sessionStorage.getItem('logged') == 'true')
  }
  getSession(){
    return(sessionStorage.getItem('token'))
  }
  
}
