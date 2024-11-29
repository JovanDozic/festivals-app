import { Component, inject, OnInit } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogActions,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle,
} from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { FestivalService } from '../../../services/festival/festival.service';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import { provideNativeDateAdapter } from '@angular/material/core';
import { MatTabsModule } from '@angular/material/tabs';
import { UserService } from '../../../services/user/user.service';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';
import { Employee } from '../../../models/festival/festival.model';
import {
  CreateUpdateUserProfileRequest,
  UpdateStaffEmailRequest,
  UpdateStaffProfileRequest,
} from '../../../models/user/user-profile-request.model';
import { forkJoin } from 'rxjs';

@Component({
  selector: 'app-edit-employee',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatInputModule,
    MatFormFieldModule,
    MatButtonModule,
    MatCardModule,
    MatDatepickerModule,
    MatGridListModule,
    MatIconModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatTabsModule,
  ],
  templateUrl: './edit-employee.component.html',
  styleUrls: ['./edit-employee.component.scss', '../../../app.component.scss'],
  providers: [provideNativeDateAdapter()],
})
export class EditEmployeeComponent implements OnInit {
  private fb = inject(FormBuilder);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<EditEmployeeComponent>);
  data: Employee = inject(MAT_DIALOG_DATA);
  private userService = inject(UserService);

  infoFormGroup: FormGroup;

  constructor() {
    this.infoFormGroup = this.fb.group({
      emailCtrl: ['', Validators.required],
      firstNameCtrl: ['', Validators.required],
      lastNameCtrl: ['', Validators.required],
      phoneNumberCtrl: ['', Validators.required],
      dateOfBirthCtrl: ['', Validators.required],
    });
  }

  ngOnInit(): void {
    if (this.data) {
      this.infoFormGroup.patchValue({
        emailCtrl: this.data.email,
        firstNameCtrl: this.data.firstName,
        lastNameCtrl: this.data.lastName,
        phoneNumberCtrl: this.data.phoneNumber,
        dateOfBirthCtrl: new Date(this.data.dateOfBirth),
      });
    }
  }

  saveChanges() {
    if (this.infoFormGroup.valid) {
      const emailRequest: UpdateStaffEmailRequest = {
        username: this.data.username,
        email: this.infoFormGroup.get('emailCtrl')?.value,
      };

      const profileRequest: UpdateStaffProfileRequest = {
        username: this.data.username,
        firstName: this.infoFormGroup.get('firstNameCtrl')?.value,
        lastName: this.infoFormGroup.get('lastNameCtrl')?.value,
        dateOfBirth: this.infoFormGroup.get('dateOfBirthCtrl')?.value,
        phoneNumber: this.infoFormGroup.get('phoneNumberCtrl')?.value,
      };

      forkJoin({
        emailUpdate: this.userService.updateStaffEmail(emailRequest),
        profileUpdate: this.userService.updateStaffProfile(profileRequest),
      }).subscribe({
        next: () => {
          this.dialogRef.close(true);
          this.snackbarService.show('Employee updated successfully');
        },
        error: (error) => {
          console.error(error);
          this.snackbarService.show('Error updating Employee');
        },
      });
    }
  }

  closeDialog() {
    this.dialogRef.close(false);
  }
}
