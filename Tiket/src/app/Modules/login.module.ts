import { NgModule } from '@angular/core';
import { MatDialogModule } from '@angular/material/dialog';
import { SocialLoginModule, AuthServiceConfig } from "angularx-social-login";
import { GoogleLoginProvider, FacebookLoginProvider } from "angularx-social-login";
 
let config = new AuthServiceConfig([
  {
    id: GoogleLoginProvider.PROVIDER_ID,
    provider: new GoogleLoginProvider("882548559741-23tpqk4qjiveh7ru4korv7v0rqslio1p.apps.googleusercontent.com")
  },
  {
    id: FacebookLoginProvider.PROVIDER_ID,
    provider: new FacebookLoginProvider("1088591984828967")
  }
]);
export function getConfig(){
  return config;
}
@NgModule({
  declarations: [ ],
  imports: [
    MatDialogModule,
    SocialLoginModule
  ],
  exports:[
    MatDialogModule,
  ],
  providers: [
    {
      provide: AuthServiceConfig,
      useFactory: getConfig
    }
  ],
})
export class LoginModule { }
