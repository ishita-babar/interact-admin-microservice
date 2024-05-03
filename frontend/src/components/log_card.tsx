import { Log } from '@/types';
import React, { Dispatch, SetStateAction, useState } from 'react';
import moment from 'moment';
import { X } from '@phosphor-icons/react';
import Toaster from '@/utils/toaster';
import deleteHandler from '@/handlers/delete_handler';
import ConfirmDelete from './confirm_delete';
import Cookies from 'js-cookie';

interface Props {
  log: Log;
  setLogs: Dispatch<SetStateAction<Log[]>>;
}

const LogCard = ({ log, setLogs }: Props) => {
  const [clickedOnDelete, setClickedOnDelete] = useState(false);

  const handleDelete = async () => {
    const toaster = Toaster.startLoad('Deleting this log...');
    const URL = `${process.env.NEXT_PUBLIC_BACKEND_URL}/logs/${log.id}`;
    const res = await deleteHandler(URL);
    if (res.statusCode === 204) {
      setLogs(prev => prev.filter(l => l.id != l.id));
      setClickedOnDelete(false);
      Toaster.stopLoad(toaster, 'Log Deleted', 1);
    } else {
      Toaster.stopLoad(toaster, 'Internal Server Error', 0);
    }
  };

  const getLogColor = () => {
    switch (log.level) {
      case 'info':
        return '#DCF9FD';
      case 'fatal':
      case 'error':
        return '#FFBABA';
      case 'debug':
        return '#DAE0FF';
      case 'warn':
      case 'warning':
        return '#FEFFDB';
      case 'success':
        return '#DBFFDD';
      default:
        return '#fff';
    }
  };

  const userRole = Cookies.get('role');

  return (
    <>
      {clickedOnDelete ? <ConfirmDelete handleDelete={handleDelete} setShow={setClickedOnDelete} /> : <></>}
      <div className="w-[95%] h-16 mx-auto border-b-[1px] border-gray-200 flex text-base text-gray-600">
        <div className="w-1/12 flex-center max-md:hidden">{moment(log.timestamp).format('HH:MM:SS')}</div>
        <div className="w-1/12 flex-center max-md:w-2/6 max-md:text-xs">
          {moment(log.timestamp).format('DD MMM YY')}
        </div>
        <div className={`${userRole == 'Manager' ? 'w-2/12' : 'w-3/12'} flex-center max-md:w-3/6 max-md:text-xs`}>
          {log.title}
        </div>
        <div className="w-1/12 flex-center max-md:w-3/6 max-md:text-xs">{log.resource}</div>
        <div className="w-1/12 flex-center max-md:w-1/6 max-md:text-xs">
          <div
            style={{ backgroundColor: getLogColor() }}
            className="w-20 rounded-lg p-1 flex-center text-sm font-medium"
          >
            {log.level}
          </div>
        </div>
        <div className="w-3/12 flex-center max-md:hidden text-xs">{log.description}</div>
        <div className="w-2/12 flex-center max-md:hidden text-sm">{log.path}</div>
        {userRole == 'Manager' && (
          <div className="w-1/12 flex-center">
            <X onClick={() => setClickedOnDelete(true)} className="cursor-pointer" size={20} />
          </div>
        )}
      </div>
    </>
  );
};

export default LogCard;
