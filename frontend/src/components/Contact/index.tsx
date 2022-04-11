import React from 'react';
import Boat from '@/images/boat.svg';
import useIsMobile from '@/hooks/useIsMobile';

import ContactForm from '../ContactForm';
import styles from './styles.module.scss';

export default function Contact() {
  const [isMobile] = useIsMobile(768);
  return (
    <div className={styles.contact}>
      {!isMobile && <Boat />}
      <div className={styles.contactFormWrapper}>
        <h1>Contact Me</h1>
        <ContactForm />
      </div>
    </div>
  );
}
