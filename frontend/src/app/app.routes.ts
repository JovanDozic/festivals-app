import { Routes } from '@angular/router';
import { AuthGuard } from './services/auth/auth.guard';
import { UnauthorizedComponent } from './shared/unauthorized/unauthorized.component';
import { AuthRedirectGuard } from './services/auth/auth-redirect.guard';

const publicRoutes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    redirectTo: 'home',
  },
  {
    path: 'home',
    loadComponent: () =>
      import('./features/home/home.component').then((c) => c.HomeComponent),
    title: 'Home',
  },
  {
    path: 'login',
    loadComponent: () =>
      import('./features/user/login/login.component').then(
        (c) => c.LoginComponent
      ),
    title: 'Login',
    canActivate: [AuthRedirectGuard],
  },
  {
    path: 'register',
    loadComponent: () =>
      import(
        './features/user/register-attendee/register-attendee.component'
      ).then((c) => c.RegisterAttendeeComponent),
    title: 'Register',
    canActivate: [AuthRedirectGuard],
  },
  {
    path: 'unauthorized',
    component: UnauthorizedComponent,
    title: 'Unauthorized',
  },
  // ...
];

const authRoutes: Routes = [
  {
    path: '',
    canActivate: [AuthGuard],
    children: [
      {
        path: 'profile',
        loadComponent: () =>
          import('./features/user/profile/profile.component').then(
            (c) => c.ProfileComponent
          ),
        title: 'Profile',
      },
      // ...
    ],
  },
];

const attendeeRoutes: Routes = [
  {
    path: '',
    canActivate: [AuthGuard],
    data: { expectedRoles: ['ATTENDEE'] },
    children: [
      // ...
    ],
  },
];

const organizerRoutes: Routes = [
  {
    path: '',
    canActivate: [AuthGuard],
    data: { expectedRoles: ['ORGANIZER'] },
    children: [
      {
        path: 'organizer/my-festivals',
        loadComponent: () =>
          import(
            './features/organizer/my-festivals/my-festivals.component'
          ).then((c) => c.MyFestivalsComponent),
        title: 'My Festivals',
      },
      {
        path: 'organizer/my-festivals/create',
        loadComponent: () =>
          import(
            './features/organizer/create-festival/create-festival.component'
          ).then((c) => c.CreateFestivalComponent),
        title: 'Create Festival',
      },
      {
        path: 'organizer/my-festivals/:id',
        loadComponent: () =>
          import('./features/organizer/festival/festival.component').then(
            (c) => c.FestivalComponent
          ),
        title: 'Festival',
      },
      {
        path: 'organizer/my-festivals/:id/employees',
        loadComponent: () =>
          import(
            './features/organizer/festival-employees/festival-employees/festival-employees.component'
          ).then((c) => c.FestivalEmployeesComponent),
        title: 'Festival Employees',
      },
      {
        path: 'organizer/my-festivals/:id/ticket-types',
        loadComponent: () =>
          import(
            './features/organizer/ticket-types/ticket-types/ticket-types.component'
          ).then((c) => c.TicketTypesComponent),
        title: 'Ticket Types',
      },
      {
        path: 'organizer/my-festivals/:id/package-addons',
        loadComponent: () =>
          import(
            './features/organizer/package-addons/package-addons/package-addons.component'
          ).then((c) => c.PackageAddonsComponent),
        title: 'Package Addons',
      },
      {
        path: 'organizer/my-festivals/:id/package-addons/general',
        loadComponent: () =>
          import(
            './features/organizer/package-addons/general-package-addons/general-package-addons.component'
          ).then((c) => c.GeneralPackageAddonsComponent),
        title: 'General Addons',
      },
      {
        path: 'organizer/my-festivals/:id/package-addons/transport',
        loadComponent: () =>
          import(
            './features/organizer/package-addons/transport-package-addons/transport-package-addons.component'
          ).then((c) => c.TransportPackageAddonsComponent),
        title: 'Travel Addons',
      },
      {
        path: 'organizer/my-festivals/:id/package-addons/camp',
        loadComponent: () =>
          import(
            './features/organizer/package-addons/camp-package-addons/camp-package-addons.component'
          ).then((c) => c.CampPackageAddonsComponent),
        title: 'Camp Addons',
      },
      // ...
    ],
  },
];

export const routes: Routes = [
  ...publicRoutes,
  ...authRoutes,
  ...attendeeRoutes,
  ...organizerRoutes,
];
