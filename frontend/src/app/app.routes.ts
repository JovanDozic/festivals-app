import { Routes } from '@angular/router';
import { AuthGuard } from './core/auth.guard';
import { UnauthorizedComponent } from './shared/unauthorized/unauthorized.component';
import { AuthRedirectGuard } from './core/auth-redirect.guard';

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
      {
        path: 'dashboard',
        loadComponent: () =>
          import('./dashboard/dashboard.component').then(
            (c) => c.DashboardComponent
          ),
        title: 'Dashboard',
      },
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
        path: 'address',
        loadComponent: () =>
          import('./address-form/address-form.component').then(
            (c) => c.AddressFormComponent
          ),
        title: 'Address',
      },
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
            './features/organizer/festival-employees/festival-employees.component'
          ).then((c) => c.FestivalEmployeesComponent),
        title: 'Festival Employees',
      },
      {
        path: 'organizer/my-festivals/:id/ticket-types',
        loadComponent: () =>
          import(
            './features/organizer/ticket-types/ticket-types.component'
          ).then((c) => c.TicketTypesComponent),
        title: 'Ticket Types',
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
