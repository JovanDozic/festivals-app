import { Component, Inject, inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatDatepickerModule } from '@angular/material/datepicker';
import {
  MatDialogTitle,
  MatDialogContent,
  MatDialogActions,
} from '@angular/material/dialog';
import { AuthService } from '../../../core/auth.service';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { UserService } from '../../../services/user/user.service';
import {
  MAT_DATE_FORMATS,
  provideNativeDateAdapter,
} from '@angular/material/core';

@Component({
  selector: 'app-change-profile-dialog',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatInputModule,
    MatDatepickerModule,
  ],
  templateUrl: './change-profile-dialog.component.html',
  styleUrl: './change-profile-dialog.component.scss',
  providers: [
    provideNativeDateAdapter(),
    // { provide: MAT_DATE_FORMATS, useValue: 'YYYY-MM-DD' },
  ],
})
export class ChangeProfileDialogComponent implements OnInit {
  private userService = inject(UserService);
  readonly dialogRef = inject(MatDialogRef<ChangeProfileDialogComponent>);
  readonly formBuilder = inject(FormBuilder);
  readonly data = inject<any>(MAT_DIALOG_DATA);

  changeProfileForm: FormGroup = this.formBuilder.group({
    firstName: ['', Validators.required],
    lastName: ['', Validators.required],
    dateOfBirth: [null, Validators.required],
    phoneNumber: [''],
  });

  ngOnInit() {
    console.log(this.data);
    if (this.data) {
      const { firstName, lastName, dateOfBirth, phoneNumber } = this.data;
      this.changeProfileForm.patchValue({
        firstName,
        lastName,
        dateOfBirth: dateOfBirth ? new Date(dateOfBirth) : null,
        phoneNumber,
      });
    }
  }

  changeProfile() {
    if (this.changeProfileForm.valid) {
      this.userService
        .updateUserProfile(this.changeProfileForm.value)
        .subscribe({
          next: () => {
            this.dialogRef.close(true);
          },
          error: (error) => {
            console.error(error);
          },
        });
    }
  }

  closeDialog() {
    this.dialogRef.close(false);
  }
}
