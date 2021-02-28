// This file can be replaced during build by using the `fileReplacements` array.
// `ng build --prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

export const environment = {
  production: false,
  auth_api_config : {
    URL: "http://localhost:30001/v1",

    basicauth_login: "admin",
    basicauth_password: "admin123"
  },
  message_api_config : {
    URL: "http://localhost:30009",

    basicauth_login: "admin",
    basicauth_password: "admin123"
  },
  abonnement_api_config : {
    URL: "http://localhost:30005",

    basicauth_login: "admin",
    basicauth_password: "admin123"
  },
  match_api_config : {
    URL: "http://localhost:30007",

    basicauth_login: "admin",
    basicauth_password: "admin123"
  }
};

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/dist/zone-error';  // Included with Angular CLI.
