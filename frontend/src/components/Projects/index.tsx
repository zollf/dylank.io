import React, { useState } from 'react';
import ReactPaginate from 'react-paginate';
import useIsMobile from '@/hooks/useIsMobile';
import { ChevronLeft, ChevronRight } from '@/images';
import { useQuery } from '@apollo/client';
import { xor } from 'lodash';

import ProjectItems from '../ProjectItems';
import Tags from '../Tags';
import query from './query.graphql';
import styles from './styles.module.scss';

const ProjectsContext = React.createContext<ProjectsContext>({
  projects: undefined,
  loading: false,
  activeFilters: [],
  setActiveFilter: () => undefined,
  clearFilters: () => undefined,
});

export default function Projects() {
  const itemsPerPage = 6;
  const [activeFilters, setActiveFilters] = useState<Array<string>>([]);
  const [pageCount, setPageCount] = useState(0);
  const [isMobile] = useIsMobile(768);
  const { data, loading, error } = useQuery<WorkQuery>(query, {
    variables: { tags: activeFilters, offset: pageCount * itemsPerPage, limit: itemsPerPage },
    fetchPolicy: 'no-cache',
  });

  const setActiveFilter = (name: string) => {
    setActiveFilters(xor([name], activeFilters));
  };
  const clearFilters = () => setActiveFilters([]);

  if (error) {
    return null;
  }

  return (
    <ProjectsContext.Provider
      value={{ projects: data?.projects, loading, clearFilters, setActiveFilter, activeFilters }}
    >
      <Tags />
      <ProjectItems />
      <div className={styles.paginationWrapper}>
        {!isMobile && (
          <ReactPaginate
            className={styles.pagination}
            breakLabel="..."
            nextLabel={<ChevronRight />}
            onPageChange={(ev) => setPageCount(ev.selected)}
            pageRangeDisplayed={5}
            forcePage={pageCount}
            pageCount={Math.ceil((data?.projects?.items_total || 6) / itemsPerPage)}
            previousLabel={<ChevronLeft />}
            pageClassName={styles.paginationPage}
            nextClassName={styles.paginationNext}
            previousClassName={styles.paginationPrev}
            activeClassName={styles.paginationActive}
            disabledClassName={styles.paginationDisabled}
          />
        )}
      </div>
    </ProjectsContext.Provider>
  );
}

export { ProjectsContext };
