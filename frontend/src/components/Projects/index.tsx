import React, { useEffect, useState } from 'react';
import useIsMobile from '@/hooks/useIsMobile';
import { ChevronLeft, ChevronRight } from '@/images';
import { useQuery } from '@apollo/client';
import { xor } from 'lodash';

import Pagination from '../Pagination';
import ProjectItems from '../ProjectItems';
import Tags from '../Tags';
import query from './query.graphql';

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
  const [data, setData] = useState<WorkQuery>();
  const {
    data: _data,
    loading,
    error,
  } = useQuery<WorkQuery>(query, {
    variables: { tags: activeFilters, offset: pageCount * itemsPerPage, limit: itemsPerPage },
    fetchPolicy: 'no-cache',
  });

  useEffect(() => {
    if (!loading) setData(_data);
  }, [loading, _data, data]);

  const setActiveFilter = (name: string) => {
    setPageCount(0);
    setActiveFilters(xor([name], activeFilters));
  };
  const clearFilters = () => setActiveFilters([]);

  if (error) {
    return null;
  }

  return (
    <ProjectsContext.Provider
      value={{ projects: data?.projects, loading: loading, clearFilters, setActiveFilter, activeFilters }}
    >
      <Tags />
      <ProjectItems />
      {!isMobile && (
        <Pagination
          nextLabel={<ChevronRight />}
          onPageChange={(ev) => setPageCount(ev.selected)}
          pageRangeDisplayed={5}
          forcePage={pageCount}
          pageCount={Math.ceil((data?.projects?.itemsTotal || 6) / itemsPerPage)}
          previousLabel={<ChevronLeft />}
        />
      )}
    </ProjectsContext.Provider>
  );
}

export { ProjectsContext };
