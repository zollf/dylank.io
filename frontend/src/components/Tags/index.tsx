import React, { useEffect, useRef, useState } from 'react';
import { useScroll } from 'react-use';
import Tag from '@/components/Tag';
import { ChevronLeft, ChevronRight } from '@/images';
import cc from 'classcat';
import styles from './styles.module.scss';

interface Props {
  totalProjects: number;
  tags: Tag[];
  setActiveFilter: (value: string) => void;
  activeFilters: string[];
}

export default function Tags({ totalProjects, tags, setActiveFilter, activeFilters }: Props) {
  const [start, setStart] = useState(true);
  const [end, setEnd] = useState(false);
  const slider = useRef<HTMLDivElement>(null);
  const { x } = useScroll(slider);
  function next() {
    if (slider.current) {
      slider.current.scrollLeft += slider.current.clientWidth;
    }
  }

  function prev() {
    if (slider.current) {
      slider.current.scrollLeft -= slider.current.clientWidth;
    }
  }

  useEffect(() => {
    if (slider.current) {
      const maxScroll = slider.current.scrollWidth - slider.current.clientWidth;
      setEnd(Math.floor(x) === Math.floor(maxScroll));
      setStart(x === 0);
    }
  }, [x, slider]);

  return (
    <div className={styles.tagsWrapper}>
      <button className={cc({ [styles.prev]: true, [styles.hidden]: start })} onClick={prev}>
        <ChevronLeft />
      </button>
      <div className={styles.tags} ref={slider}>
        <Tag onClick={() => setActiveFilter('all')} active={activeFilters.includes('all')}>
          All ({totalProjects})
        </Tag>
        {tags.map((tag) => (
          <Tag onClick={() => setActiveFilter(tag.id)} active={activeFilters.includes(tag.id)}>
            {tag.title} ({tag.count})
          </Tag>
        ))}
      </div>
      <button className={cc({ [styles.next]: true, [styles.hidden]: end })} onClick={next}>
        <ChevronRight />
      </button>
    </div>
  );
}
