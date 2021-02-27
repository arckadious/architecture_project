import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';

import { MatchsPage } from './matchs.page';

describe('MatchsPage', () => {
  let component: MatchsPage;
  let fixture: ComponentFixture<MatchsPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MatchsPage ],
      imports: [IonicModule.forRoot()]
    }).compileComponents();

    fixture = TestBed.createComponent(MatchsPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
