import MockDate from 'mockdate';
import fetchMock from 'jest-fetch-mock';

MockDate.set('2000-01-01');

class ResizeObserver {
  observe() {
    // do nothing
  }
  unobserve() {
    // do nothing
  }
  disconnect() {
    // do nothing
  }
}

jest.mock('three/examples/jsm/loaders/GLTFLoader', () => ({}));

window.ResizeObserver = ResizeObserver;

fetchMock.enableMocks();
