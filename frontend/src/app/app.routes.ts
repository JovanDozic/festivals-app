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
        (c) => c.LoginComponent,
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
            (c) => c.ProfileComponent,
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
        path: 'festivals',
        loadComponent: () =>
          import(
            './features/attendee/all-festivals/all-festivals.component'
          ).then((c) => c.AllFestivalsComponent),
        title: 'Festivals',
      },
      {
        path: 'festivals/:id',
        loadComponent: () =>
          import('./features/attendee/festival/festival.component').then(
            (c) => c.FestivalComponent,
          ),
        title: 'Festival',
      },
      {
        path: 'festivals/:id/store/ticket',
        loadComponent: () =>
          import(
            './features/attendee/store/store-ticket/store-ticket.component'
          ).then((c) => c.StoreTicketComponent),
        title: 'Festival Ticket Store',
      },
      {
        path: 'festivals/:id/store/package',
        loadComponent: () =>
          import(
            './features/attendee/store/store-package/store-package.component'
          ).then((c) => c.StorePackageComponent),
        title: 'Festival Package Store',
      },
      {
        path: 'my-orders',
        loadComponent: () =>
          import(
            './features/attendee/order/my-orders/my-orders.component'
          ).then((c) => c.MyOrdersComponent),
        title: 'My Orders',
      },
      {
        path: 'my-orders/:id',
        loadComponent: () =>
          import('./features/attendee/order/order/order.component').then(
            (c) => c.OrderComponent,
          ),
        title: 'View Order',
      },
      {
        path: 'my-bracelets',
        loadComponent: () =>
          import(
            './features/attendee/bracelet/my-bracelets/my-bracelets.component'
          ).then((c) => c.MyBraceletsComponent),
        title: 'My Bracelets',
      },
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
            (c) => c.FestivalComponent,
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
      {
        path: 'organizer/my-festivals/:id/orders',
        loadComponent: () =>
          import(
            './features/organizer/order/festival-orders/festival-orders.component'
          ).then((c) => c.FestivalOrdersComponent),
        title: 'Orders',
      },
      {
        path: 'organizer/my-festivals/:id/orders/:orderId',
        loadComponent: () =>
          import('./features/organizer/order/order/order.component').then(
            (c) => c.OrderComponent,
          ),
        title: 'Orders',
      },
      // ...
    ],
  },
];

const employeeRoutes: Routes = [
  {
    path: '',
    canActivate: [AuthGuard],
    data: { expectedRoles: ['EMPLOYEE'] },
    children: [
      {
        path: 'employee/my-festivals',
        loadComponent: () =>
          import(
            './features/employee/my-festivals/my-festivals.component'
          ).then((c) => c.MyFestivalsComponent),
        title: 'My Festivals',
      },
      {
        path: 'employee/my-festivals/:id',
        loadComponent: () =>
          import('./features/employee/festival/festival.component').then(
            (c) => c.FestivalComponent,
          ),
        title: 'Festival',
      },
      {
        path: 'employee/my-festivals/:id/employees',
        loadComponent: () =>
          import(
            './features/employee/festival-employees/festival-employees.component'
          ).then((c) => c.FestivalEmployeesComponent),
        title: 'Festival Employees',
      },
      {
        path: 'employee/my-festivals/:id/ticket-types',
        loadComponent: () =>
          import(
            './features/employee/ticket-types/ticket-types/ticket-types.component'
          ).then((c) => c.TicketTypesComponent),
        title: 'Ticket Types',
      },
      {
        path: 'employee/my-festivals/:id/package-addons',
        loadComponent: () =>
          import(
            './features/employee/package-addons/package-addons/package-addons.component'
          ).then((c) => c.PackageAddonsComponent),
        title: 'Package Addons',
      },
      {
        path: 'employee/my-festivals/:id/package-addons/general',
        loadComponent: () =>
          import(
            './features/employee/package-addons/general-package-addons/general-package-addons.component'
          ).then((c) => c.GeneralPackageAddonsComponent),
        title: 'General Addons',
      },
      {
        path: 'employee/my-festivals/:id/package-addons/transport',
        loadComponent: () =>
          import(
            './features/employee/package-addons/transport-package-addons/transport-package-addons.component'
          ).then((c) => c.TransportPackageAddonsComponent),
        title: 'Travel Addons',
      },
      {
        path: 'employee/my-festivals/:id/package-addons/camp',
        loadComponent: () =>
          import(
            './features/employee/package-addons/camp-package-addons/camp-package-addons.component'
          ).then((c) => c.CampPackageAddonsComponent),
        title: 'Camp Addons',
      },
      {
        path: 'employee/my-festivals/:id/orders',
        loadComponent: () =>
          import(
            './features/employee/order/festival-orders/festival-orders.component'
          ).then((c) => c.FestivalOrdersComponent),
        title: 'Orders',
      },
      {
        path: 'employee/my-festivals/:id/orders/:orderId',
        loadComponent: () =>
          import('./features/employee/order/order/order.component').then(
            (c) => c.OrderComponent,
          ),
        title: 'Orders',
      },
    ],
  },
];

const adminRoutes: Routes = [
  {
    path: '',
    canActivate: [AuthGuard],
    data: { expectedRoles: ['ADMINISTRATOR'] },
    children: [
      {
        path: 'admin/users',
        loadComponent: () =>
          import('./features/admin/all-users/all-users.component').then(
            (c) => c.AllAccountsComponent,
          ),
        title: 'All Users',
      },
      {
        path: 'admin/users/:id',
        loadComponent: () =>
          import('./features/admin/user/user.component').then(
            (c) => c.AccountComponent,
          ),
        title: 'Users',
      },
      {
        path: 'admin/logs',
        loadComponent: () =>
          import('./features/admin/all-logs/all-logs.component').then(
            (c) => c.AllLogsComponent,
          ),
        title: 'Logs',
      },
    ],
  },
];

export const routes: Routes = [
  ...publicRoutes,
  ...authRoutes,
  ...attendeeRoutes,
  ...organizerRoutes,
  ...employeeRoutes,
  ...adminRoutes,
];
