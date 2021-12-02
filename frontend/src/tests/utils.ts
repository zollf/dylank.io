import { act } from 'react-dom/test-utils';

const setMobile = async (width = 400) => {
  act(() => {
    Object.defineProperty(window, 'innerWidth', { writable: true, configurable: true, value: width });
    window.dispatchEvent(new Event('resize'));
  });
  jest.mock('@/hooks/useIsMobile', () => {
    [true];
  });
};

export { setMobile };
