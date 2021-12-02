import React from 'react';
import cc from 'classcat';

import styles from './styles.module.scss';

interface BaseProps {
  children?: React.ReactNode;
  theme: 'blue' | 'white';
  size: 'small' | 'large';
  full?: boolean;
}

interface AnchorProps
  extends Omit<React.DetailedHTMLProps<React.AnchorHTMLAttributes<HTMLAnchorElement>, HTMLAnchorElement>, 'type'> {
  href: string;
  onClick?: never;
  type?: never;
}

interface ButtonProps
  extends Omit<React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>, 'type'> {
  onClick: (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
  href?: never;
  type?: string;
}

interface SubmitButtonProps
  extends Omit<React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>, 'type'> {
  type?: 'submit';
  onClick?: never;
  href?: never;
}

type Props = (AnchorProps | ButtonProps | SubmitButtonProps) & BaseProps;

const Button = ({ children, type, theme, size, full, ...props }: Props) => {
  const Wrapper = props.href ? 'a' : 'button';

  let classNames = cc({
    [styles.button]: true,
    [styles.full]: full,
  });

  if (theme && styles[theme]) {
    classNames = cc([classNames, styles[theme]]);
  }

  if (size && styles[size]) {
    classNames = cc([classNames, styles[size]]);
  }

  return (
    <Wrapper type={props.href ? undefined : type} role="button" tabIndex={0} className={classNames} {...(props as any)}>
      {children}
    </Wrapper>
  );
};

export default Button;
