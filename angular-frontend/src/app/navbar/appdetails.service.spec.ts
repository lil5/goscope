import { TestBed } from '@angular/core/testing';

import { AppdetailsService } from './appdetails.service';

describe('AppdetailsService', () => {
  let service: AppdetailsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AppdetailsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
