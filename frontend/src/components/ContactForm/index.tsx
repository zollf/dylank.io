import React, { useCallback, useState } from 'react';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';
import { FormControl, FormLabel, Input, Stack, Textarea } from '@chakra-ui/react';

import Button from '../Button';
import styles from './styles.module.scss';

interface ContactFormFields {
  name: string;
  email: string;
  message: string;
}

export default function ContactForm() {
  const [sent, setIsSent] = useState(false);
  const onSubmit = useCallback(async () => {
    setIsSent(true);
    return;
  }, []);

  return (
    <div className={styles.contactForm}>
      <Formik<ContactFormFields> initialValues={initialValues} onSubmit={onSubmit} validationSchema={validationSchema}>
        {({ handleSubmit, errors, touched }) => (
          <Form onSubmit={handleSubmit}>
            <Stack spacing={6}>
              <FormControl isInvalid={!!errors.name && touched.name} isDisabled={sent}>
                <FormLabel {...defaultProps} htmlFor="name">
                  Name
                </FormLabel>
                <Field as={Input} id="name" name="name" placeholder="John Smith" />
              </FormControl>
              <FormControl isInvalid={!!errors.email && touched.email} isDisabled={sent}>
                <FormLabel {...defaultProps} htmlFor="email">
                  Email
                </FormLabel>
                <Field as={Input} id="email" name="email" type="email" placeholder="your@email.com" />
              </FormControl>
              <FormControl isInvalid={!!errors.message && touched.message} isDisabled={sent}>
                <FormLabel {...defaultProps} htmlFor="message">
                  Message
                </FormLabel>
                <Field
                  as={Textarea}
                  id="message"
                  name="message"
                  placeholder="Hey, I would like to message you about..."
                  resize="none"
                />
              </FormControl>
              {!sent ? (
                <Button type="submit" theme="blue" size="large">
                  Send
                </Button>
              ) : (
                <strong>Test Sent</strong>
              )}
            </Stack>
          </Form>
        )}
      </Formik>
    </div>
  );
}

const initialValues = {
  name: '',
  email: '',
  message: '',
};

const defaultProps = {
  mb: '8px',
  fontWeight: 'semibold',
};

const validationSchema = Yup.object().shape({
  name: Yup.string().required(),
  email: Yup.string().email().required(),
  message: Yup.string().required(),
});
