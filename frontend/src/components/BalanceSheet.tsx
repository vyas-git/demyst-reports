import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { API_END_POINT } from '../config/app';

interface Cell {
  Value: string;
  Attributes?: Array<{ Value: string; Id: string }>;
}

interface Row {
  RowType: string;
  Cells?: Cell[];
  Title?: string;
  Rows?: Row[];
}

interface Report {
  ReportID: string;
  ReportName: string;
  ReportType: string;
  ReportTitles: string[];
  ReportDate: string;
  UpdatedDateUTC: string;
  Rows: Row[];
}

const BalanceSheet: React.FC = () => {
  const [report, setReport] = useState<Report | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    axios.get(`${API_END_POINT}/accounts/reports/balance-sheet`)
      .then(response => {
        setReport(response.data.Reports[0]);
      })
      .catch(error => {
        setError('Error fetching data');
      });
  }, []);

  if (error) {
    return <div className="text-red-500">{error}</div>;
  }

  if (!report) {
    return <div className="text-gray-500">Loading...</div>;
  }

  if (report.Rows.length === 0) {
    return <div className="text-gray-500">No records</div>;
  }

  const renderRows = (rows: Row[]) => {
    return rows.map((row, index) => {
      if (row.RowType === 'Header') {
        return (
          <tr key={index} className="bg-gray-100 text-black">
            {row.Cells?.map((cell, cellIndex) => (
              <th key={cellIndex} className="p-2 border">{cell.Value}</th>
            ))}
          </tr>
        );
      } else if (row.RowType === 'Section') {
        return (
          <React.Fragment key={index}>
            <tr className="bg-gray-100 text-black">
              <td colSpan={3} className="p-2 border text-left"><strong>{row.Title}</strong></td>
            </tr>
            {row.Rows && renderRows(row.Rows)}
          </React.Fragment>
        );
      } else if (row.RowType === 'SummaryRow') {
        return (
          <tr key={index} className="bg-gray-200 text-black">
            {row.Cells?.map((cell, cellIndex) => (
              <td key={cellIndex} className={`p-2 border ${cellIndex === 0 ? 'font-bold text-left' : 'font-bold text-right'}`}>{cell.Value}</td>
            ))}
          </tr>
        );
      } else {
        return (
          <tr key={index} className="bg-gray-100 text-black">
            {row.Cells?.map((cell, cellIndex) => (
              <td key={cellIndex} className="p-2 border text-right">{cell.Value}</td>
            ))}
          </tr>
        );
      }
    });
  };

  return (
    <div className="container mx-auto p-4">
      <div className="bg-gray-600 border border-gray-700 rounded-md p-4 mb-4">
        <h1 className="text-3xl font-bold text-gray-100 mb-2">{report.ReportTitles[0]}</h1>
        <h2 className="text-2xl text-gray-300 mb-1">{report.ReportTitles[1]}</h2>
        <h3 className="text-xl text-gray-400 mb-4">{report.ReportTitles[2]}</h3>
      </div>
      <table className="w-full border-collapse">
        <thead>
          {renderRows(report.Rows.filter(row => row.RowType === 'Header'))}
        </thead>
        <tbody>
          {renderRows(report.Rows.filter(row => row.RowType !== 'Header'))}
        </tbody>
      </table>
    </div>
  );
};

export default BalanceSheet;
