import { useEffect, useState } from 'react';
import { useWindowSize } from 'react-use';

export default function useIsMobile(...cutoffParams: number[]) {
  const { width: x } = useWindowSize();
  const [isMobile, setIsMobile] = useState<Array<boolean>>([]);
  const [cutoffs] = useState(cutoffParams || [768]);

  useEffect(() => {
    const checked = cutoffs.map((co) => x <= co);
    setIsMobile(checked);
  }, [x, cutoffs]);

  return isMobile;
}
