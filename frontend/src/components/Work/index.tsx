import React, { useEffect, useState } from 'react';
import ReactPaginate from 'react-paginate';
import Project from '@/components/Project';
import useIsMobile from '@/hooks/useIsMobile';
import { ChevronLeft, ChevronRight } from '@/images';
import { useQuery } from '@apollo/client';
import { xor } from 'lodash';

import Tags from '../Tags';
import query from './query.graphql';
import styles from './styles.module.scss';

interface WorkQuery {
  tags: Array<Tag>;
  projects: Array<Project>;
}

export default function Work() {
  const itemsPerPage = 6;
  const [activeFilters, setActiveFilters] = useState<Array<string>>(['all']);

  const [itemOffset, setItemOffset] = useState(0);
  const [pageCount, setPageCount] = useState(0);
  const [currentItems, setCurrentItems] = useState<Array<Project>>([]);

  const [isMobile] = useIsMobile(768);
  const { data, loading, error } = useQuery<WorkQuery>(query);

  const setActiveFilter = (name: string) => {
    if (name === 'all') {
      setActiveFilters(['all']);
      return;
    }

    let currentFilters = activeFilters;
    if (currentFilters[0] === 'all') currentFilters = [];
    setActiveFilters(xor([name], currentFilters));
  };

  const handlePageClick = (event: { selected: number }) => {
    const newOffset = (event.selected * itemsPerPage) % (data?.projects.length || 1);
    setItemOffset(newOffset);
  };

  useEffect(() => {
    if (!activeFilters.length) {
      setActiveFilters(['all']);
    }
  }, [activeFilters]);

  useEffect(() => {
    if (!!data?.projects?.length) {
      if (isMobile) {
        setCurrentItems(data.projects);
      } else {
        const endOffset = itemOffset + itemsPerPage;
        setCurrentItems(data.projects.slice(itemOffset, endOffset));
        setPageCount(Math.ceil(data.projects.length / itemsPerPage));
      }
    }
  }, [itemOffset, itemsPerPage, data?.projects, isMobile]);

  if (!data || loading || error) {
    return null;
  }

  return (
    <div className={styles.work}>
      <div className={styles.top}>
        <h1>Work</h1>
        {!isMobile && (
          <p>
            Check my commercial and non commercial projects. If you have any questions feel free to ask me for more
            information.
          </p>
        )}
      </div>
      <Tags
        totalProjects={data.projects.length}
        tags={data.tags}
        setActiveFilter={setActiveFilter}
        activeFilters={activeFilters}
      />
      <div className={styles.projects}>
        {currentItems.map((project) => (
          <Project project={project} />
        ))}
      </div>
      {!isMobile && (
        <ReactPaginate
          className={styles.pagination}
          breakLabel="..."
          nextLabel={<ChevronRight />}
          onPageChange={handlePageClick}
          pageRangeDisplayed={5}
          pageCount={pageCount}
          previousLabel={<ChevronLeft />}
          pageClassName={styles.paginationPage}
          nextClassName={styles.paginationNext}
          previousClassName={styles.paginationPrev}
          activeClassName={styles.paginationActive}
          disabledClassName={styles.paginationDisabled}
        />
      )}
    </div>
  );
}
