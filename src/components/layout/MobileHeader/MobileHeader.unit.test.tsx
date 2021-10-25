import React from 'react';
import UserEvent from '@testing-library/user-event';
import { act, render, waitFor, waitForElementToBeRemoved } from '@testing-library/react';
import { setMobile } from '@/tests/utils';

import MobileHeader from './';

jest.useFakeTimers();

it('matches its snapshot', () => {
  const { baseElement } = render(<MobileHeader />);
  expect(baseElement).toMatchSnapshot();
});

it('matches its mobile snapshot', () => {
  setMobile();
  const { baseElement } = render(<MobileHeader />);
  expect(baseElement).toMatchSnapshot();
});

it('opens and close mobile menu correctly', async () => {
  setMobile();
  const { getByTestId } = render(<MobileHeader />);
  UserEvent.click(getByTestId('hbtn'));
  await waitFor(() => expect(getByTestId('backing')).toBeInTheDocument());

  act(() => {
    jest.runAllTimers();
  });

  UserEvent.click(getByTestId('hbtn'));
  await waitForElementToBeRemoved(() => getByTestId('backing'));
});
