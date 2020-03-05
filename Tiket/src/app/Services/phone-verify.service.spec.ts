import { TestBed } from '@angular/core/testing';

import { PhoneVerifyService } from './phone-verify.service';

describe('PhoneVerifyService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: PhoneVerifyService = TestBed.get(PhoneVerifyService);
    expect(service).toBeTruthy();
  });
});
