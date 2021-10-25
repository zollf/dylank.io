import React, { useCallback, useState } from 'react';
import CloseIcon from '@/images/icon/close.svg';
import MenuIcon from '@/images/icon/burger.svg';

import MobileMenu from '../MobileMenu';
import styles from './styles.module.scss';

const MobileHeader = () => {
  const [active, setActive] = useState(false);
  const [delayClosed, setDelayedClosed] = useState(true);

  const toggle = useCallback(() => {
    active ? close() : open();
  }, [active]);

  const open = useCallback(() => {
    setDelayedClosed(false);
    setTimeout(() => setActive(true), 50);
  }, [setDelayedClosed, setActive]);

  const close = useCallback(() => {
    setActive(false);
    setTimeout(() => setDelayedClosed(true), 200);
  }, [setActive, setDelayedClosed]);

  return (
    <>
      <header className={styles.mobileHeader}>
        <div className={styles.logo}>d.io</div>
        <button type="button" className={styles.menu} onClick={toggle} data-active={active} data-testid="hbtn">
          {active ? <CloseIcon /> : <MenuIcon />}
        </button>
        <MobileMenu active={active} />
      </header>
      {!delayClosed && (
        <button type="button" className={styles.backing} onClick={close} data-active={active} data-testid="backing" />
      )}
    </>
  );
};

export default MobileHeader;