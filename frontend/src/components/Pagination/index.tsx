import React from 'react';
import ReactPaginate, { ReactPaginateProps } from 'react-paginate';

import styles from './styles.module.scss';

export default function Pagination(props: ReactPaginateProps) {
  return (
    <div className={styles.paginationWrapper}>
      <ReactPaginate
        className={styles.pagination}
        breakLabel="..."
        pageClassName={styles.paginationPage}
        nextClassName={styles.paginationNext}
        previousClassName={styles.paginationPrev}
        activeClassName={styles.paginationActive}
        disabledClassName={styles.paginationDisabled}
        {...props}
      />
    </div>
  );
}
